<template>
  <div class="dashboard-container">
    <fileManager></fileManager>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import fileManager from "../../components/fileManager"

export default {
  name: 'Dashboard',
  components:{
    fileManager
  },
  computed: {
    ...mapGetters([
      'name'
    ])
  }
}
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
