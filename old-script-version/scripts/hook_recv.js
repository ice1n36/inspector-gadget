Interceptor.attach(Module.findExportByName('libc.so', 'recv'), {
  onEnter: function (args) {
    send('recv(' +
      'socket=' + args[0] +
      ', buffer=' + args[1] +
      ', length=' + args[2] +
      ', flags=' + args[3] +
    ')');
  },
  onLeave: function (retvat) {
  }
});
