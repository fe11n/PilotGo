package dao

import (
	"openeluer.org/PilotGo/PilotGo/pkg/app/server/model"
	"openeluer.org/PilotGo/PilotGo/pkg/logger"
	"openeluer.org/PilotGo/PilotGo/pkg/mysqlmanager"
)

func IsExistName(name string) bool {
	var batch model.Batch
	mysqlmanager.DB.Where("name=?", name).Find(&batch)
	return batch.ID != 0
}
func GetBatchID(name string) uint {
	var machine model.Batch
	mysqlmanager.DB.Where("name=?", name).Find(&machine)
	return machine.ID
}
func GetmachineBatch(uuid string) string {
	var machine model.MachineNode
	mysqlmanager.DB.Where("machine_uuid=?", uuid).Find(&machine)
	return machine.BatchInfo
}

func UpdatemachineBatch(s string, b string) {
	var machineInfo model.MachineNode
	machine := model.MachineNode{
		BatchInfo: b,
	}
	mysqlmanager.DB.Model(&machineInfo).Where("machine_uuid=?", s).Update(&machine)
	logger.Info("%+v", machine)
}

// func Batchinfo() []model.Batch {
// 	var batch []model.Batch
// 	a := mysqlmanager.DB.Find(&batch)
// 	logger.Info("%+v", a)
// 	return batch
// }

// func UpdatemachineBatch(m model.MachineNode) {

// }
