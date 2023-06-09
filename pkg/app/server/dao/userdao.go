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
 * LastEditTime: 2022-03-16 15:25:41
 * Description: 用户模块相关数据获取
 ******************************************************************************/
package dao

import (
	"fmt"
	"strconv"
	"strings"

	"openeuler.org/PilotGo/PilotGo/pkg/app/server/model"
	"openeuler.org/PilotGo/PilotGo/pkg/global"
)

// 获取所有的用户角色
func AllUserRole() []model.UserRole {
	var role []model.UserRole
	global.PILOTGO_DB.Find(&role)
	return role
}

// 邮箱账户是否存在
func IsEmailExist(email string) bool {
	var user model.User
	global.PILOTGO_DB.Where("email=?", email).Find(&user)
	return user.ID != 0
}

// 查询数据库中账号密码、用户部门、部门ID、用户类型、用户角色
func UserPassword(email string) (s1, s2, s3 string, i1, i2 int) {
	var user model.User
	global.PILOTGO_DB.Where("email=?", email).Find(&user)
	return user.Password, user.DepartName, user.RoleID, user.DepartSecond, user.UserType
}

// 查询某用户信息
func UserInfo(email string) model.User {
	var user model.User
	global.PILOTGO_DB.Where("email=?", email).Find(&user)
	return user
}

// 查询所有的用户
func UserAll() ([]model.ReturnUser, int) {
	var users []model.User
	var redisUser []model.ReturnUser

	// 先从redis缓存中读取
	// data, err := redismanager.Get("users", &redisUser)
	// if err == nil {
	// 	resByre, _ := json.Marshal(data)
	// 	json.Unmarshal(resByre, &redisUser)
	// 	logger.Debug("%+v", "从缓存中读取")
	// 	return redisUser
	// } else {
	global.PILOTGO_DB.Order("id desc").Find(&users)
	totals := len(users)
	for _, user := range users {
		var roles []string
		// 查找角色
		roleids := user.RoleID
		roleId := strings.Split(roleids, ",")
		for _, id := range roleId {
			userRole := model.UserRole{}
			i, _ := strconv.Atoi(id)
			global.PILOTGO_DB.Where("id = ?", i).Find(&userRole)
			role := userRole.Role
			roles = append(roles, role)
		}
		u := model.ReturnUser{
			ID:           user.ID,
			DepartFirst:  user.DepartFirst,
			DepartSecond: user.DepartSecond,
			DepartName:   user.DepartName,
			Username:     user.Username,
			Phone:        user.Phone,
			Email:        user.Email,
			UserType:     user.UserType,
			Roles:        roles,
		}
		redisUser = append(redisUser, u)
	}
	// redismanager.Set("users", &redisUser)
	return redisUser, totals
	// }
}

// 根据用户邮箱模糊查询
func UserSearch(email string) ([]model.ReturnUser, int) {
	var users []model.User
	var redisUser []model.ReturnUser

	global.PILOTGO_DB.Order("id desc").Where("email LIKE ?", "%"+email+"%").Find(&users)
	totals := len(users)
	for _, user := range users {
		var roles []string
		// 查找角色
		roleids := user.RoleID
		roleId := strings.Split(roleids, ",")
		for _, id := range roleId {
			userRole := model.UserRole{}
			i, _ := strconv.Atoi(id)
			global.PILOTGO_DB.Where("id = ?", i).Find(&userRole)
			role := userRole.Role
			roles = append(roles, role)
		}
		u := model.ReturnUser{
			ID:           user.ID,
			DepartFirst:  user.DepartFirst,
			DepartSecond: user.DepartSecond,
			DepartName:   user.DepartName,
			Username:     user.Username,
			Phone:        user.Phone,
			Email:        user.Email,
			UserType:     user.UserType,
			Roles:        roles,
		}
		redisUser = append(redisUser, u)
	}
	return redisUser, totals
}

// 重置密码
func ResetPassword(email string) (model.User, error) {
	var user model.User
	global.PILOTGO_DB.Where("email=?", email).Find(&user)
	if user.ID != 0 {
		global.PILOTGO_DB.Model(&user).Where("email=?", email).Update("password", "123456")
		return user, nil
	} else {
		return user, fmt.Errorf("无此用户")
	}
}

// 删除用户
func DeleteUser(email string) {
	var user model.User
	global.PILOTGO_DB.Where("email=?", email).Unscoped().Delete(user)
}

// 修改用户的部门信息
func UpdateUserDepart(email, departName string, Pid, id int) {
	var user model.User
	u := model.User{
		DepartFirst:  Pid,
		DepartSecond: id,
		DepartName:   departName,
	}
	global.PILOTGO_DB.Model(&user).Where("email=?", email).Updates(&u)
}

// 添加用户
func AddUser(u model.User) {
	global.PILOTGO_DB.Save(&u)
}

// 修改手机号
func UpdateUserPhone(email, phone string) {
	var user model.User
	global.PILOTGO_DB.Model(&user).Where("email=?", email).Update("phone", phone)
}

func DelUser(deptId int) {
	var user model.User
	global.PILOTGO_DB.Where("depart_second=?", deptId).Unscoped().Delete(user)
}
