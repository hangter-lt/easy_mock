import { createMemoryHistory, createRouter } from 'vue-router'

import ApiList from '../components/ApiList.vue'
import ApiEdit from '../components/ApiEdit.vue'
import ApiInfo from '../components/ApiInfo.vue'
import ApiRealtime from '../components/ApiRealtime.vue'


const routes = [
  { path: '/', component: ApiRealtime },
  { path: '/apilist', component: ApiList },
  { path: '/apiedit/:id', component: ApiEdit },
  { path: '/apiinfo/:id', component: ApiInfo },  
]

const router = createRouter({
  history: createMemoryHistory(),
  routes,
})

export default router