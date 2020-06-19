<template>
  <el-drawer
    title="控制台"
    :visible.sync="visible"
    size="60%"
    direction="ltr">
    <el-container>
      <el-main style="max-height: 80vh;height: 80vh">
        <el-input
          type="textarea"
          :autosize="{ minRows: 10, maxRows: 40}"
          placeholder="请输入内容"
          v-model="outPut">
        </el-input>
      </el-main>
      <el-footer>
        <div style="margin-top: 15px;">
          <el-input placeholder="请输入内容" v-model="inputContent">
            <template slot="append">
              <el-button type="primary"
              @click="websocketsend">执行</el-button>
            </template>
          </el-input>
        </div>
      </el-footer>
    </el-container>
  </el-drawer>
</template>

<script>
  import ElButton from "element-ui/packages/button/src/button";

  export default {
    components: {ElButton},
    name: "index",
    props: ["path"],
    data() {
      return {
        websock: null,
        visible: false,
        inputContent :"",
        outPut:""
      }
    },
    created() {
      this.initWebSocket();
    },
    destroyed() {
      this.closeWebSocket()
    },
    methods: {
      initWebSocket() { //初始化weosocket
        const wsuri = "ws://" + window.location.host.split(":").slice(0,1) + ":8895/file-ws/cmd";
        this.websock = new WebSocket(wsuri);
        this.websock.onmessage = this.websocketonmessage;
        // this.websock.onopen = this.websocketonopen;
        this.websock.onerror = this.websocketonerror;
        this.websock.onclose = this.websocketclose;
      },
      closeWebSocket() {
        if (this.websock) {
          this.websock.close()
          this.websock = null
        }
      },
      websocketonerror(err) {//连接建立失败重连
        this.initWebSocket();
      },
      websocketonmessage(e) { //数据接收
        try {
          var obj = JSON.parse(e.data)
          this.outPut += obj.body
        }catch {
          this.outPut += "err: " +e.data
        }
      },
      websocketsend() {
        let data  = {
          "path":this.path,
          "cmd":this.inputContent
        }
        this.outPut += "cmd : "+ this.inputContent + "\n"
        this.websock.send(JSON.stringify(data));
      },
      websocketclose(e) {  //关闭
        this.initWebSocket()
        console.log('断开连接', e);
      },
    },
  }
</script>

<style scoped>

</style>
