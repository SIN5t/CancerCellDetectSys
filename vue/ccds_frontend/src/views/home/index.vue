<template>
  <div>
    <label for="file-upload">选择文件: </label>
    <input id="file-upload" type="file" ref="fileInput" @change="handleFileChange">
    <button @click="uploadFile">上传文件</button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      file: null
    }
  },
  methods: {
    handleFileChange(event) {
      // 获取选择的文件
      this.file = event.target.files[0]
    },
    async uploadFile() {
      if (!this.file) {
        alert('请先选择文件')
        return
      }
      const formData = new FormData()
      formData.append('file_content', this.file)
      try {
        const response = await axios.post('http://localhost:8888/ccds/upload', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })
        console.log(response.data)
        alert('文件上传成功')
      } catch (error) {
        console.error(error)
        alert('文件上传失败')
      }
    }
  }
}
</script>
