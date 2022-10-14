package nakamacluster

import (
	"fmt"
	"testing"

	"github.com/doublemo/nakama-cluster/api"
	addr "github.com/hashicorp/go-sockaddr"
	"github.com/serialx/hashring"
)

func TestIp(t *testing.T) {
	fmt.Println(addr.GetPrivateIP())
	fmt.Println(addr.GetPublicIP())

	v, err := addr.NewIPAddr("189.2.3.56:7895")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v.NetIP(), v.IPPort())
	// fmt.Println(PublicIP())
	// fmt.Println(LocalIP())

	weights := make(map[string]int)
	weights["A"] = 1
	weights["B"] = 2
	//weights["C"] = 1

	ring := hashring.NewWithWeights(weights)
	// key := ring.GenKey("A")
	fmt.Println(ring.GetNodePos("A-0"))
	fmt.Println(ring.GetNodePos("B-0"))
	fmt.Println(ring.GetNodePos("B-1"))
	fmt.Println(ring.GetNodePos("C-0"))
	fmt.Println(ring.GetNode("A-0"))
	fmt.Println(ring.GetNode("B-0"))
	fmt.Println(ring.GetNode("B-1"))
	fmt.Println(ring.GetNode("f11e1c3c-0db2-48bc-a16d-c29dbf51dc19"))
}

func TestMessageQueue(t *testing.T) {
	a := NewMessageQueue(10)
	a.Push(&api.Envelope{Id: 1})
	a.Push(&api.Envelope{Id: 2})
	a.Push(&api.Envelope{Id: 3})

	t.Log("count - >", a.Len())
	fmt.Println(a.PopAll())
	a.Push(&api.Envelope{Id: 4})
	t.Log("count - >", a.Len(), a.Pop())
}
