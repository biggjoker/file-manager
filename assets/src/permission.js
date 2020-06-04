import router from './router'
import store from './store'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import getPageTitle from '@/utils/get-page-title'

NProgress.configure({showSpinner: false}) // NProgress Configuration

var bFirst = true

router.beforeEach(async (to, from, next) => {
  // start progress bar
  NProgress.start()

  // set page title
  document.title = getPageTitle(to.meta.title)
  // if (bFirst) {
  //   bFirst = false
  //   const accessRoutes = await store.dispatch('permission/generateRoutes', 'admin')
  //   router.addRoutes(accessRoutes)
  // }
  next()
})

router.afterEach(() => {
  // finish progress bar
  NProgress.done()
})
