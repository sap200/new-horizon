<!-- AlertModal.vue -->
<template>
  <div v-if="isVisible" class="modal-overlay">
    <div class="modal">
      <p id="m-id">{{ message }}</p>
      <br><br>
      <hr>
      <Button class="fdb-btn" @click="close" >Close</Button>
    </div>
  </div>
</template>

<script>
export default {
  name: "AlertModal",
  props: {
    isVisible: Boolean,
    message: String
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
  background-color: #f0f0f0;
  padding: 20px;
  border-radius: 5px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
  word-wrap: break-word; /* Use break-word to allow breaking long words */
  max-width: 600px; /* Set a max-width if needed */
}

#m-id {
    color: black;
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

</style>
