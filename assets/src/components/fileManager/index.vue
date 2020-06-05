<template>
  <el-card class="box-card">
    <div slot="header" class="clearfix">
      <div style="margin-bottom: 20px">
        <el-upload
          class="inline-block"
          :on-change="handleChange"
          :show-file-list="false"
          :data="actionData"
          :limit="999"
          :multiple="true"
          name="file"
          :auto-upload="true"
          action="/file-manager/dir/upfile"
          :file-list="fileList">
          <el-button size="small" type="primary">点击上传</el-button>
        </el-upload>
        <el-button size="small"
                   type="primary"
                   @click="createDic">创建文件夹
        </el-button>
        <!--<el-button size="small"-->
        <!--v-if="path === '/'"-->
        <!--type="primary">创建文本</el-button>-->
      </div>
    </div>
    <div slot="header" class="clearfix">
      <span>当前路径 </span>
      <span class="path">{{path}}</span>
      <el-button
        icon="el-icon-back"
        @click="backPage"
        circle
        v-if="path !== '/'">
      </el-button>
    </div>
    <div>
      <folder v-for="(item,index) in filesCurr"
              :key="index"
              @reload="reloadTable"
              :info="item">
      </folder>
    </div>
  </el-card>
</template>

<script>
  import folder from "./components/folder"
  import {getDic,
  createDic} from "../../api/file"
  import ElButton from "element-ui/packages/button/src/button";

  export default {
    name: "index",
    components: {
      ElButton,
      folder
    },
    data() {
      return {
        files: [],
        path: "/",
        fileList: [],
      }
    },
    methods: {
      reloadTable() {
        let path = this.$route.query.path
        if (!path) {
          path = "/"
        }
        this.path = path
        getDic(path).then(res => {
          this.files = res
        })
      },
      handleChange(file, fileList) {
        this.reloadTable()
      },
      backPage() {
        if (this.path !== "/") {
          this.$router.push({
            query: {
              path: this.path.split("/").slice(0, -1).join("/")
            }
          })
        }
      },
      createDic(){
        this.$prompt('请输入名称', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
        }).then(({ value }) => {
          let path = this.path
          if(path === "/"){
            path += value
          }else {
            path +=  "/" +value
          }
          createDic(path).then(res =>{
            this.$message.success("创建成功")
            this.reloadTable()
          })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '取消输入'
          })
        })
      }
    },
    mounted() {
      this.reloadTable()
    },
    computed: {
      filesCurr() {
        let tt = this.files.filter(t => {
          if (t.name.length <= this.path.length) {
            return false
          }
          let head = t.name.slice(0, this.path.length)
          let tail = t.name.slice(this.path.length)
          if(this.path === "/"){
            tail = "/" + tail
          }

          return head === this.path &&
            tail[0] === "/" &&
            tail.indexOf("/", 1) === -1
        })
        return tt
      },
      actionData() {
        return {
          path: this.path
        }
      }
    },
    watch: {
      '$route'(to, from) {
        this.reloadTable()
      }
    }
  }
</script>

<style scoped>
  .path {
    color: #67C23A;
    font-size: 15px;
    margin-left: 50px;
    margin-right: 20px;
  }

  .inline-block {
    display: inline-block;
    margin-right: 12px;
  }
</style>
