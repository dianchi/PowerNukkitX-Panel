package src

import "strconv"

func Para(javapath string, Xms int, Xmx int, path string, corename string) string {
	var in string
	XMS := strconv.Itoa(Xms) + "m"
	XMX := strconv.Itoa(Xmx) + "m"
	in += javapath + " "
	in += JVMDPara("file.encoding=" + EncodingFormat)
	in += JVMDPara("jansi.passthrough=true")
	in += JVMDPara("terminal.ansi=true")
	in += JVMPara("UnlockExperimentalVMOptions")
	in += JVMPara("UseG1GC")
	in += JVMPara("UseStringDeduplication")
	in += JVMPara("EnableJVMCI")
	in += "-Xms" + XMS + " "
	in += "-Xmx" + XMX + " "
	in += "--module-path=" + path + "/libs/graal-sdk-22.2.0.jar;" + path + "/libs/truffle-api-22.2.0.jar;" + " "
	in += "--upgrade-module-path=;" + " "
	in += JVMParaWithPara("java.base/java.lang")
	in += JVMParaWithPara("java.base/java.io")
	in += "-cp" + " " + path + "/" + corename + ";" + path + "/libs/*" + " "
	in += "cn.nukkit.Nukkit"
	return in
}

func JVMParaWithPara(in string) string {
	out := "--add-opens" + " " + in + "=ALL-UNNAMED" + " "
	return out
}

func JVMPara(in string) string {
	out := "-XX:+" + in + " "
	return out
}

func JVMDPara(in string) string {
	out := "-D" + in + " "
	return out
}
