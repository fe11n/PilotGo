<!-- 
  Copyright (c) KylinSoft Co., Ltd.2021-2022. All rights reserved.
  PilotGo is licensed under the Mulan PSL v2.
  You can use this software accodring to the terms and conditions of the Mulan PSL v2.
  You may obtain a copy of Mulan PSL v2 at:
      http://license.coscl.org.cn/MulanPSL2
  THIS SOFTWARE IS PROVIDED ON AN 'AS IS' BASIS, WITHOUT WARRANTIES OF ANY KIND, 
  EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
  See the Mulan PSL v2 for more details.
  Author: zhaozhenfang
  Date: 2022-02-25 16:33:45
  LastEditTime: 2022-06-24 16:39:58
  Description: provide agent log manager of pilotgo
 -->
<template>
 <div style="width: 100%;height:100%">
   <ky-table
        class="cluster-table"
        ref="table"
        :showSelect="showSelect"
        :getData="getLogs"
        :searchData="searchData"
      >
        <template v-slot:table_search>
          <div>日志列表</div>
        </template>
        <template v-slot:table_action>
        </template>
        <template v-slot:table>
          <el-table-column prop="type" label="日志名称">
          </el-table-column>
          <el-table-column prop="userName" label="创建者"> 
          </el-table-column>
          <el-table-column prop="departName" label="部门"> 
          </el-table-column>
          <el-table-column label="进度">
            <template slot-scope="scope">
              <el-progress style="width: 100%" 
                type="line" 
                text-inside
                :stroke-width="strokeW"
                :format="format"
                :percentage="(scope.row.status.split(',')[2] === '1.00' || scope.row.status.split(',')[2] === '0.00') ? 100 : scope.row.status.split(',')[2] * 100 || 0" 
                :status="scope.row.status.split(',')[2] === '0.00' ? 'exception' : 
                    scope.row.status.split(',')[2] === '1.00' ? 'success' : 'warning' ">
              </el-progress>
            </template> 
          </el-table-column>
          <el-table-column prop="created_at" label="创建时间" sortable>
            <template slot-scope="scope">
              <span>{{scope.row.created_at | dateFormat}}</span>
            </template>
          </el-table-column>
          <el-table-column prop="operation" label="详情">
            <template slot-scope="scope">
              <el-button
                size="mini"
                type="primary"
                plain
                @click="handleDetail(scope.row)">
                查看
              </el-button>
            </template>
          </el-table-column>
        </template>
      </ky-table>
      <el-dialog 
        :title="title"
        :before-close="handleClose" 
        :visible.sync="display" 
        :width="dialogWidth"
      >
        <log-detail class="detail" v-if="type === 'detail'" :log="log" @click="handleClose"></log-detail>
      </el-dialog>
 </div>
</template>
<script>
import kyTable from "@/components/KyTable";
import LogDetail from "./form/detail.vue"
import { getLogs, deleteLog } from "@/request/log";
export default {
  name: "Log",
  components: {
    kyTable,
    LogDetail,
  },
  data() {
    return {
      dialogWidth: '70%',
      display: false,
      title: '',
      type: '',
      log: {},
      strokeW: 15,
      searchData: {
        departId: this.$store.getters.UserDepartId || '',
      },
      showSelect: true,
    }
  },
  methods: {
    getLogs,
    refresh(){
      this.$refs.table.handleSearch();
    },
    handleClose() {
      this.display = false;
      this.title = "";
      this.type = "";
    },
    handleDetail(row) {
      // 查看日志详情
      this.display = true;
      this.title = "日志详情";
      this.type = "detail";
      this.log = row;
    },
    format(value) {
      if(value) {
        return `${value}%`;
      }
    }
  },
  filters: {
    dateFormat: function(value) {
      let date = new Date(value);
      let y = date.getFullYear();
      let MM = date.getMonth() + 1;
      MM = MM < 10 ? "0" + MM : MM;
      let d = date.getDate();
      d = d < 10 ? "0" + d : d;
      let h = date.getHours();
      h = h < 10 ? "0" + h : h;
      let m = date.getMinutes();
      m = m < 10 ? "0" + m : m;
      let s = date.getSeconds();
      s = s < 10 ? "0" + s : s;
      return y + "-" + MM + "-" + d + " " + h + ":" + m;
    }
  },
}
</script>
<style scoped>
  .detail {
    height: 60vh;
  }
</style>
