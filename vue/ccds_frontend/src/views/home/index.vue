<!--
 * @Descripttion: 
 * @version: 
 * @Date: 2021-04-20 11:06:21
 * @LastEditors: huzhushan@126.com
 * @LastEditTime: 2022-09-24 18:18:43
 * @Author: huzhushan@126.com
 * @HomePage: https://huzhushan.gitee.io/vue3-element-admin
 * @Github: https://github.com/huzhushan/vue3-element-admin
 * @Donate: https://huzhushan.gitee.io/vue3-element-admin/donate/
-->

<template>
  <div class="file-upload">
    <input
      type="file"
      ref="fileInput"
      @change="handleFileChange"
      class="file-input"
    />
    <button @click="uploadFile" class="upload-button">上传</button>
    <p v-if="selectedFile" class="file-info">
      文件名: {{ selectedFile.name }}
      <br />
      文件大小: {{ formatFileSize(selectedFile.size) }}
    </p>
  </div>
</template>

<style>
.file-upload {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.file-input {
  margin-bottom: 10px;
}

.upload-button {
  padding: 10px 20px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.file-info {
  margin-top: 10px;
  text-align: center;
}
</style>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      selectedFile: null,
    }
  },
  methods: {
    handleFileChange() {
      this.selectedFile = this.$refs.fileInput.files[0]
    },
    uploadFile() {
      const file = this.$refs.fileInput.files[0]

      // 构建 FormData 对象
      const formData = new FormData()
      formData.append('file_content', file)

      // 发送 POST 请求到后端接口
      axios
        .post('/ccds/upload', formData)
        .then(response => {
          // 处理上传成功的逻辑
        })
        .catch(error => {
          // 处理上传失败的逻辑
        })
    },
    formatFileSize(size) {
      if (size === 0) return '0 Bytes'
      const k = 1024
      const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']
      const i = Math.floor(Math.log(size) / Math.log(k))
      return parseFloat((size / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
    },
  },
}
</script>
