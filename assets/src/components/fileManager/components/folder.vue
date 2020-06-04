<template>
  <div class="wrap"
       @mouseover="changeShowInfo(true)"
       @click="clickItem()"
       @mouseout="changeShowInfo(false)">
    <div class="head">
      <el-popover
        placement="top"
        trigger="hover"
        width="160">
        <template>
          <el-button type="primary"
                     size="mini"
                     @click="dealMenuFunc('rename')"
                     icon="el-icon-edit"></el-button>
          <el-button type="danger"
                     size="mini"
                     @click="dealMenuFunc('delete')"
                     icon="el-icon-delete"></el-button>
        </template>
        <el-tag type="success"
                slot="reference"
                size="mini"
                icon="el-icon-check"
                round>
          详情
        </el-tag>
      </el-popover>
    </div>
    <div class="body">
      <i class="el-icon-folder-opened" v-if="info.is_dir"
         style="color: #ffd46d;"></i>
      <i class="el-icon-notebook-1" v-else
         style="color: #5bc9f5"></i>
    </div>
    <div :title="info.name" class="footer">
      {{ info.name.split("/").slice(-1)[0] }}
    </div>

    <el-dialog title="文本信息"
               :visible.sync="text.visible"
               :close-on-click-modal="false"
               width="800px"
               :append-to-body="true">
        <span slot="title">
            <span>文本信息</span>
        </span>
      <el-form label-width="80px">
        <el-form-item>
          <el-button type="success"
                     style="float: right"
                     size="mini"
                     @click="saveFileContent">
            保存
          </el-button>
        </el-form-item>
        <el-form-item
          @keydown.ctrl.83.native="handleEditorConfirm">
          <el-input
            type="textarea"
            :autosize="{ minRows: 5, maxRows: 30}"
            placeholder="请输入内容"
            v-model="text.content">
          </el-input>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
  import {
    readFile,
    renameDic,
    delDic,
    saveFile
  } from "../../../api/file"

  export default {
    name: "folder",
    props: ["info"],
    data() {
      return {
        ifShow: false,
        text: {
          visible: false,
          content: ""
        }
      }
    },
    methods: {
      changeShowInfo(val) {
        setTimeout(() => {
          this.ifShow = val
        }, 0)
      },
      clickItem() {
        if (this.info.is_dir) {
          this.$router.push({
            query: {
              path: this.info.name
            }
          })
        } else {
          this.showFile()
        }
      },
      showFile() {
        if (this.checkFileSize(this.info.size)) {
          readFile(this.info.name).then(res => {
            this.text.content = res
            this.text.visible = true

          })
        } else {
          this.$message.info("文件大于5M 不支持在线查看")
        }
      },
      checkFileType(name) {
        let type = name.split(".").slice(-1)[0]
        console.log(type)
        return this.StaticType.textType.findIndex(t => {
          return t === type
        }) !== -1
      },
      checkFileSize(size) {
        return size <= 5 * 1024 * 1024
      },
      dealMenuFunc(type) {
        if (type === "rename") {
          let nameOld = this.info.name.split("/").slice(-1)[0]
          this.$prompt('请输入 ' + nameOld + ' 新的名称', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
          }).then(({value}) => {
            let pathOld = this.info.name
            let pathNew = this.info.name.split("/").slice(0, -1).join("/") + "/" + value
            renameDic(pathNew, pathOld).then(res => {
              this.$message.success("修改成功")
              this.$emit("reload")
            })
          }).catch(() => {
            this.$message({
              type: 'info',
              message: '取消输入'
            })
          })
        } else if (type === "delete") {
          this.$confirm('确认删除服务 删除后不可恢复')
            .then(() => {
              delDic(this.info.name, true).then(res => {
                this.$message.success("操作成功")
                this.$emit("reload")
              })
            }).catch(() => {
          })
        }
      },
      saveFileContent() {
        saveFile(this.info.name,
          this.text.content).then(res => {
          this.$message.success("保存成功")
          this.text.visible = false
          this.text.content = ""
        })
      },
      handleEditorConfirm(e) {
        e.stopPropagation()
        e.preventDefault()
        this.saveFileContent()
      }
    },
  }
</script>

<style scoped>
  .wrap {
    width: 100px;
    height: 100px;
    border: 5px;
    margin: 5px;
    float: left;
    cursor: pointer;
  }

  .wrap:hover {
    background: #f1f5fa;
  }

  .head {
    float: right;
    margin-top: 5px;
    margin-right: 5px;
  }

  .body {
    margin-top: 20px;
    font-size: 40px;
    text-align: center;
  }

  .footer {
    text-align: center;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    font-size: 10px;
  }


</style>
