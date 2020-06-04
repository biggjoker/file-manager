const okResult = {
  'meta': {
    'success': true,
    'message': 'ok',
    'code': 0,
  },
  'data': {
    roles: ['admin'],
    introduction: 'I am a super administrator',
    avatar: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',
    name: 'Super Admin',
  },
}

export default {
  login() {
    return okResult
  },
  getInfo() {
    return okResult
  }
}
