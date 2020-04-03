import React from "react";
import {protocols, Terminal, WebTTY} from "@/pages/Dashboard/webtty";
import {Xterm} from "@/pages/Dashboard/xterm";
import {ConnectionFactory} from "@/pages/Dashboard/websocket";

export interface KnifeTerminalProps {
  onRef?: Function,
}

export default class KnifeTerminal extends React.PureComponent<KnifeTerminalProps> {
  term: Terminal | undefined;
  state = {
    width: 600,
    height: 500
  };

  constructor(props: any) {
    super(props);
    if (this.props.onRef) {
      this.props.onRef(this)
    }
  }

  resize() {
    this.term?.resize()
  }

  componentDidMount(): void {
    var gotty_auth_token = "";

    const elem = document.getElementById("terminal")

    if (elem !== null) {
      //let term: Terminal;
      this.term = new Xterm(elem);
      const httpsEnabled = window.location.protocol == "https:";
      const url = (httpsEnabled ? 'wss://' : 'ws://') + window.location.host + '/ws/tty/';
      const args = window.location.search;
      console.log("url:" + url)
      const factory = new ConnectionFactory(url, protocols);
      const wt = new WebTTY(this.term, factory, args, gotty_auth_token);
      const closer = wt.open();
      const tmpTerm = this.term
      window.addEventListener("unload", () => {
        closer();
        tmpTerm.close();
      });
    } else {
      console.log("ele is null")
    }
  }

  render() {
    return (
      <div id="terminal" style={{
        width: "100%",
        height:"calc(100% - 20px)"
      }}/>
    )
  }
}
