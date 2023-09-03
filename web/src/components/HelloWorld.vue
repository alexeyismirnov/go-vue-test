<template>
  <v-container class="align-center text-center">
    <v-row class="align-center justify-center">
      <v-col>
      <span class="display-1 grey--text"> {{ message }} </span>
      </v-col>
    </v-row>

    <v-row class="align-center justify-center">
      <v-col cols="auto">
        <v-btn flat color="primary" class="display-1" @click="openDialog">Add message</v-btn>
      </v-col>

      <v-col cols="auto">
        <v-btn flat color="error" class="display-1" @click="reset">Clear message</v-btn>
      </v-col>
    </v-row>
  </v-container>

  <Popup  @response="(msg) => message = msg" />

</template>

<script>
  import Popup from '@/components/Popup.vue'
  import axios from 'axios'

  export default {
    data() {
      return {
        message: ''
      }
    },

    components: { Popup },

    methods: {
      openDialog() {
        this.emitter.emit("dialog", true);
      },

      reset() {
        this.message = ''
      }
    },

    mounted() {
      axios.get("http://127.0.0.1:8080/api/v1/getMessage").then(response => this.message = response.data)
      // this.message = localStorage.getItem('message') || 'Please type a message'
    },

    watch: {
      message(newVal) {
        axios.post("http://127.0.0.1:8080/api/v1/setMessage", newVal)
        // localStorage.setItem('message', newVal)
      }
    },

  }

</script>
