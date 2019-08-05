package controller

import (
	"fmt"
	"github.com/micro/go-micro/util/log"
	"github.com/dexinq/sdn_global/api/controller/proto"
	"github.com/dexinq/sdn_global/utils/cmd"
	"gopkg.in/mgo.v2/bson"
)

type Chassis struct {
	Hostname string `json:"hostname" bson:"hostname"`
}

type LoadBalance struct {
	Name     string              `json:"name" bson:"name"`
	Protocal string              `json:"protocal" bson:"protocal"`
	Rules    []map[string]string `json:"rules" bson:"rules"`
}

type LSPort struct {
	Name string `json:"name" bson:"name"`
	Ip   string `json:"ip" bson:"ip"`
	Mac  string `json:"mac"  bson:"mac"`
}

type LogicalSwitch struct {
	Name string         `json:"name" bson:"name"`
	Port []*LSPort      `json:"port" bson:"port"`
	LB   []*LoadBalance `json:"loadbalancer" bson:"loadbalancer"`
}

type LRPort struct {
	Name string `json:"name" bson:"name"`
	Ip   string `json:"ip"  bson:"ip"`
	Mac  string `json:"mac" bson:"mac"`
}

type LogicalRouter struct {
	Name string         `json:"name" bson:"name"`
	Port []*LRPort      `json:"port" bson:"port"`
	LB   []*LoadBalance `json:"loadbalancer" bson:"loadbalancer"`
}

type Ovn struct {
	Name string             `json:"name" bson:"name"`
	Oc   *controller.OvnController `json:"config" bson:"config"`
	Ls   []*LogicalSwitch   `json:"logicalswitch" bson:"logicalswitch"`
	Lr   []*LogicalRouter   `json:"logicalrouter" bson:"logicalrouter"`
	Ch   []*Chassis         `json:"chassis" bson:"chassis"`
}

func NewOvn(cfg *controller.OvnController) *Ovn {
	return &Ovn{Oc: cfg}
}

func (o *Ovn) SyncData() error {
	if err := o.syncNb(); err != nil {
		return err
	}
	if err := o.syncSb(); err != nil {
		return err
	}
	if err := o.storeCache(); err != nil {
		return err
	}
	return nil
}

func (o *Ovn) syncNb() error {
	if err := o.getLogicalSwitch(); err != nil {
		return err
	}
	if err := o.getLogicalRouter(); err != nil {
		return err
	}
	return nil
}

func (o *Ovn) syncSb() error {
	if err := o.getChassis(); err != nil {
		return err
	}
	return nil
}

func (o *Ovn) getChassis() error {
	return nil
}

func (o *Ovn) generateUrl() string {
	return fmt.Sprintf("tcp:%s:%s", o.Oc.Nbip, o.Oc.Nbport)
}

func (o *Ovn) storeCache() error {
	// TODO Add mongodb connection
	_, err := bson.Marshal(o)
	if err != nil {
		log.Fatal("marshal bson err %v", err)
	}

	return nil
}

func (o *Ovn) getLogicalSwitch() error {
	stdout, stderr, err := cmd.RunOVNSbctl(o.generateUrl(), "list", "logical_switch")
	if err != nil {
		log.Fatalf("err occur during get_logical_switch %v", err)
	}
	log.Log("stdout is %s", stdout)
	log.Log("stderr is %s", stderr)
	return nil
}

func (o *Ovn) addLogicalSwitch() error {
	return nil
}

func (o *Ovn) modifyLogicalSwitch() error {
	return nil
}

func (o *Ovn) getLogicalRouter() error {
	return nil
}

func (o *Ovn) addLogicalRouter() error {
	return nil
}

func (o *Ovn) modifyLogicalRouter() error {
	return nil
}
