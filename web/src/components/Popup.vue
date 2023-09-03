<template>
  <v-dialog max-width="600px" v-model="dialog">
    <v-card >
      <v-card-title>
        Add message
      </v-card-title>

      <v-card-subtitle>
        This message will be saved!
      </v-card-subtitle>

      <v-card-text>
        <v-form class="px-3" ref="form" @submit.prevent="submit">
          <div class="d-flex flex-column">
            <v-text-field label="Message" v-model="message" prepend-icon="mdi-pencil"></v-text-field>

            <v-alert type="error" density="compact" variant="outlined" class="mb-4"  v-if="alerta">
              Type some text
            </v-alert>
          </div>
        </v-form>
      </v-card-text>

      <v-card-actions>
        <v-btn color="orange" @click="submit">Add</v-btn>
      </v-card-actions>

    </v-card>
  </v-dialog>
</template>

<script>

export default {
  data() {
    return {
      dialog: false,
      alerta: false,
      message: ''
    }
  },

  emits: ['response'],

  methods: {
    submit() {
      if(this.message.length > 0) {
        this.$emit('response', this.message)

        this.alerta = false;
        this.dialog = false;
        this.message = '';

      } else {
        this.alerta = true;
      }
    }
  },

  mounted() {
    this.emitter.on("dialog", (data) => {
      this.dialog = data;
    });
  }
}

</script>
