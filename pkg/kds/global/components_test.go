package global_test

import (
	"context"
	"fmt"
	mesh_proto "github.com/Kong/kuma/api/mesh/v1alpha1"
	"github.com/Kong/kuma/pkg/core/resources/apis/mesh"
	"github.com/Kong/kuma/pkg/core/resources/model"
	"github.com/Kong/kuma/pkg/core/resources/store"
	"github.com/Kong/kuma/pkg/kds/global"
	sync_store "github.com/Kong/kuma/pkg/kds/store"
	"github.com/Kong/kuma/pkg/plugins/resources/memory"
	"github.com/Kong/kuma/pkg/test/grpc"
	kds_setup "github.com/Kong/kuma/pkg/test/kds/setup"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sync"
)

var _ = Describe("Global Sync", func() {

	var localStores []store.ResourceStore
	var globalStore sync_store.SyncResourceStore
	var closeFunc func()

	BeforeEach(func() {
		serverStreams := []*grpc.MockServerStream{}
		wg := &sync.WaitGroup{}
		localStores = []store.ResourceStore{}
		for i := 0; i < 2; i++ {
			wg.Add(1)
			localStore := memory.NewStore()
			serverStream := kds_setup.StartServer(localStore, wg)
			serverStreams = append(serverStreams, serverStream)
			localStores = append(localStores, localStore)
		}

		globalStore = sync_store.NewSyncResourceStore(memory.NewStore())
		stopCh := make(chan struct{})
		clientStreams := []*grpc.MockClientStream{}
		for _, ss := range serverStreams {
			clientStreams = append(clientStreams, ss.ClientStream())
		}
		kds_setup.StartClient(clientStreams, []model.ResourceType{mesh.DataplaneType}, stopCh, global.Callbacks(globalStore))

		closeFunc = func() {
			close(stopCh)
			for _, ss := range serverStreams {
				close(ss.RecvCh)
			}
			wg.Wait()
		}
	})

	dataplaneFunc := func(cluster, service string) mesh_proto.Dataplane {
		return mesh_proto.Dataplane{
			Networking: &mesh_proto.Dataplane_Networking{
				Address: "192.168.0.1",
				Inbound: []*mesh_proto.Dataplane_Networking_Inbound{{
					Port: 1212,
					Tags: map[string]string{
						mesh_proto.ClusterTag: cluster,
						mesh_proto.ServiceTag: service,
					},
				}},
				Outbound: []*mesh_proto.Dataplane_Networking_Outbound{
					{
						Tags: map[string]string{
							mesh_proto.ServiceTag:  "web",
							mesh_proto.ProtocolTag: "http",
						},
					},
				},
			},
		}
	}

	It("should add resource to global store after adding it to global", func() {
		for i := 0; i < 10; i++ {
			dp := dataplaneFunc("kuma-cluster-1", fmt.Sprintf("service-1-%d", i))
			err := localStores[0].Create(context.Background(), &mesh.DataplaneResource{Spec: dp}, store.CreateByKey(fmt.Sprintf("dp-1-%d", i), "mesh-1"))
			Expect(err).ToNot(HaveOccurred())
		}

		Eventually(func() int {
			actual := mesh.DataplaneResourceList{}
			err := globalStore.List(context.Background(), &actual)
			Expect(err).ToNot(HaveOccurred())
			return len(actual.Items)
		}, "5s", "100ms").Should(Equal(10))

		closeFunc()
	})

	It("should sync resources independently for each Local", func() {
		for i := 0; i < 10; i++ {
			dp := dataplaneFunc("kuma-cluster-1", fmt.Sprintf("service-1-%d", i))
			err := localStores[0].Create(context.Background(), &mesh.DataplaneResource{Spec: dp}, store.CreateByKey(fmt.Sprintf("dp-1-%d", i), "mesh-1"))
			Expect(err).ToNot(HaveOccurred())
		}

		for i := 0; i < 10; i++ {
			dp := dataplaneFunc("kuma-cluster-2", fmt.Sprintf("service-2-%d", i))
			err := localStores[1].Create(context.Background(), &mesh.DataplaneResource{Spec: dp}, store.CreateByKey(fmt.Sprintf("dp-2-%d", i), "mesh-1"))
			Expect(err).ToNot(HaveOccurred())
		}

		Eventually(func() int {
			actual := mesh.DataplaneResourceList{}
			err := globalStore.List(context.Background(), &actual)
			Expect(err).ToNot(HaveOccurred())
			return len(actual.Items)
		}, "3s", "100ms").Should(Equal(20))

		err := localStores[0].Delete(context.Background(), &mesh.DataplaneResource{}, store.DeleteByKey("dp-1-0", "mesh-1"))
		Expect(err).ToNot(HaveOccurred())

		err = localStores[0].Delete(context.Background(), &mesh.DataplaneResource{}, store.DeleteByKey("dp-1-1", "mesh-1"))
		Expect(err).ToNot(HaveOccurred())

		Eventually(func() int {
			actual := mesh.DataplaneResourceList{}
			err := globalStore.List(context.Background(), &actual)
			Expect(err).ToNot(HaveOccurred())
			return len(actual.Items)
		}, "5s", "100ms").Should(Equal(18))

		closeFunc()
	})

})