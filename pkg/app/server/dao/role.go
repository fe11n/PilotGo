/******************************************************************************
 * Copyright (c) KylinSoft Co., Ltd.2021-2022. All rights reserved.
 * PilotGo is licensed under the Mulan PSL v2.
 * You can use this software accodring to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *     http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN 'AS IS' BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * Author: zhanghan
 * Date: 2021-01-24 15:08:08
 * LastEditTime: 2022-04-27 17:25:41
 * Description: 角色模块相关数据获取
 ******************************************************************************/
package dao

import (
	"fmt"
	"strconv"
	"strings"

	"openeluer.org/PilotGo/PilotGo/pkg/app/server/model"
	"openeluer.org/PilotGo/PilotGo/pkg/dbmanager/mysqlmanager"
)

// 根据角色名称返回角色id和用户类型
func GetRoleIdAndUserType(role string) (roleId string, user_type int) {
	var Role model.UserRole
	mysqlmanager.DB.Where("role = ?", role).Find(&Role)
	roleID := strconv.Itoa(Role.ID)
	var userType int
	if Role.ID > model.OrdinaryUserRoleId {
		userType = model.OtherUserType
	} else {
		userType = Role.ID - 1
	}
	return roleID, userType
}

// 根绝id获取该角色的所有信息
func RoleIdToGetAllInfo(roleid int) model.UserRole {
	var role model.UserRole
	mysqlmanager.DB.Where("id=?", roleid).Find(&role)
	return role
}

// 登录用户的权限按钮
func PermissionButtons(button string) []string {
	IDs := strings.Split(button, ",")
	var buttons []string
	for _, id := range IDs {
		var SubButton model.RoleButton
		i, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		mysqlmanager.DB.Where("id = ?", i).Find(&SubButton)
		button := SubButton.Button
		buttons = append(buttons, button)
	}
	return buttons
}

//获取所有的用户角色
func GetAllRoles() ([]model.ReturnUserRole, int) {
	var roles []model.UserRole
	var getRole []model.ReturnUserRole
	mysqlmanager.DB.Order("id desc").Find(&roles)
	total := len(roles)

	for _, role := range roles {
		var buts []string

		if len(role.ButtonID) == 0 {
			r := model.ReturnUserRole{
				ID:          role.ID,
				Role:        role.Role,
				Type:        role.Type,
				Description: role.Description,
				Menus:       role.Menus,
				Buttons:     []string{},
			}
			getRole = append(getRole, r)
			continue
		} else {
			buttonss := strings.Split(role.ButtonID, ",")

			for _, button := range buttonss {
				var but model.RoleButton
				i, _ := strconv.Atoi(button)
				mysqlmanager.DB.Where("id=?", i).Find(&but)
				buts = append(buts, but.Button)
			}
			r := model.ReturnUserRole{
				ID:          role.ID,
				Role:        role.Role,
				Type:        role.Type,
				Description: role.Description,
				Menus:       role.Menus,
				Buttons:     buts,
			}
			getRole = append(getRole, r)
		}
	}
	return getRole, total
}

// 新增角色
func AddRole(r model.UserRole) error {
	role := r.Role
	if len(role) == 0 {
		return fmt.Errorf("用户角色不能为空")
	}
	userRole := model.UserRole{
		Role:        role,
		Type:        r.Type,
		Description: r.Description,
	}
	mysqlmanager.DB.Save(&userRole)
	return nil
}

// 是否有用户绑定某角色
func IsUserBindingRole(roleId int) bool {
	var users []model.User
	mysqlmanager.DB.Find(&users)
	for _, user := range users {
		id := user.RoleID
		if find := strings.Contains(id, strconv.Itoa(roleId)); find {
			return true
		}
	}
	return false
}

// 删除用户角色
func DeleteRole(roleId int) {
	var UserRole model.UserRole
	mysqlmanager.DB.Where("id = ?", roleId).Unscoped().Delete(UserRole)
}

// 修改角色名称
func UpdateRoleName(roleId int, name string) {
	var UserRole model.UserRole
	mysqlmanager.DB.Model(&UserRole).Where("id = ?", roleId).Update("role", name)
}

// 修改角色描述
func UpdateRoleDescription(roleId int, desc string) {
	var UserRole model.UserRole
	mysqlmanager.DB.Model(&UserRole).Where("id = ?", roleId).Update("description", desc)

}

// 变更用户角色权限
func UpdateRolePermission(permission model.RolePermissionChange) model.UserRole {
	var userRole model.UserRole
	// 数组切片转为string
	menus := strings.Replace(strings.Trim(fmt.Sprint(permission.Menus), "[]"), " ", ",", -1)
	buttonId := strings.Replace(strings.Trim(fmt.Sprint(permission.ButtonId), "[]"), " ", ",", -1)

	r := model.UserRole{
		Menus:    menus,
		ButtonID: buttonId,
	}
	mysqlmanager.DB.Model(&userRole).Where("id = ?", permission.RoleID).Update(&r)
	return userRole
}
