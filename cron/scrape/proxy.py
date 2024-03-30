from subprocess import Popen, DEVNULL

class Proxy:
    def __init__(self, id: int, change_ip_after: int = 5) -> None:
        self.id = id
        self.start_port = 9060
        self.config_file = f"/tmp/torrc.{id}"
        self.data_dir = f"/tmp/tor.{id}"
        self.process: Popen | None = None

    def __get_socks_port(self):
        return self.start_port + (self.id * 2)
    
    def socks_connection(self) -> str:
        return f"socks5://localhost:{self.__get_socks_port()}"

    def __create_tor_config(self):
        change_ip_after = 5
        contents = f"""
SocksPort {self.__get_socks_port()}
ControlPort {self.__get_socks_port() + 1}
DataDirectory {self.data_dir} 
MaxCircuitDirtiness {change_ip_after}
"""
        with open(self.config_file, "w") as f:
            f.write(contents)

    def run(self) -> str:
        self.__create_tor_config()
        self.process = Popen(['tor', '-f', self.config_file], stdout=DEVNULL)
        return self.socks_connection()

    def kill(self):
        self.process.terminate()
        self.process.wait()

    def reset(self):
        self.kill()
        self.run()

