/**
 * main.js
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

import App from './App.vue'
import { createApp } from 'vue'
import { registerPlugins } from '@/plugins'
import mitt from 'mitt'

const app = createApp(App)
registerPlugins(app)

const emitter = mitt()
app.config.globalProperties.emitter = emitter

app.mount('#app')
