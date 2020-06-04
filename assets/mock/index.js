import Mock from 'mockjs'

import user from './user'

Mock.mock(/\/user\/login/, 'post', user.login)
Mock.mock(/\/user\/info/, 'get', user.getInfo)
