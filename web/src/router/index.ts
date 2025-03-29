import { createMemoryHistory, createRouter } from 'vue-router'

import HelloWorld from '../components/HelloWorld.vue'
import ApiList from '../components/ApiList.vue'
import ApiEdit from '../components/ApiEdit.vue'
import ApiInfo from '../components/ApiInfo.vue'


const routes = [
  { path: '/', component: HelloWorld },
  { path: '/apilist', component: ApiList },
  { path: '/apiedit/:id', component: ApiEdit },
  { path: '/apiinfo/:id', component: ApiInfo },  
]

const router = createRouter({
  history: createMemoryHistory(),
  routes,
})

export default router