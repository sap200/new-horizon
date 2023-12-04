<!-- AlertModal.vue -->
<template>
  <div v-if="isVisible" class="modal-overlay">
    <div class="modal">
      <p id="m-id">{{ message }}</p>
      <br><br>
                <img src="../assets/txn.jpg" v-if="noCloseButton"/>

        <div class="rectangular-box" v-if="noCloseButton">
          <div class="loader"></div>
      </div>

      <hr  v-if="!noCloseButton">
      <Button class="fdb-btn" @click="close" v-if="!noCloseButton" >Close</Button>
    </div>
  </div>
</template>

<script>
export default {
  name: "AlertModal",
  props: {
    isVisible: Boolean,
    message: String,
    noCloseButton: Boolean
  },
  methods: {
    close() {
      this.$emit('close');
    }
  }
}
</script>

<style>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.6);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal {
  background-image: linear-gradient(-45deg,pink, black,black, black, skyblue);
  padding: 20px;
  border-radius: 5px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
  word-wrap: break-word; /* Use break-word to allow breaking long words */
  max-width: 600px; /* Set a max-width if needed */
}

#m-id {
    color: orange;
}

.fdb-btn {
    padding: 5px;
    border-radius: 5px;
    color:black;
    background-image: linear-gradient(to right, skyblue, aliceblue);
    cursor: pointer;
    width: 100%;
    font-family: monospace;
}

.fdb-btn:hover:enabled { 
    background-image: linear-gradient(to left, skyblue, aliceblue);
    box-shadow: 1px 1px 10px pink; /* Add the box shadow effect on hover */

}
.rectangular-box {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 600px; /* Width of the rectangular box */
  height: 100px; /* Height of the rectangular box */
  margin: auto; /* Center the box horizontally */
  position: relative;
}

.loader {
  border: 5px solid #f3f3f3; /* Light grey background */
  border-top: 5px solid #3498db; /* Blue color for the moving part */
  border-radius: 50%;
  width: 40px;
  height: 40px;
  animation: spin 2s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

img {
  width: 200px;
  height: auto;
}

</style>
