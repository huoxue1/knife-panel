import React from "react";
import {protocols, Terminal, WebTTY} from "@/pages/Dashboard/webtty";
import {Xterm} from "@/pages/Dashboard/xterm";
import {ConnectionFactory} from "@/pages/Dashboard/websocket";


export default class KnifeTerminal extends React.PureComponent {
  componentDidMount(): void {
    var gotty_auth_token = "";

    const elem = document.getElementById("terminal")

    if (elem !== null) {
      let term: Terminal;
      term = new Xterm(elem);
      const httpsEnabled = window.location.protocol == "https:";
      const url = (httpsEnabled ? 'wss://' : 'ws://') + window.location.host + '/ws/tty/';
      const args = window.location.search;
      console.log("url:" + url)
      const factory = new ConnectionFactory(url, protocols);
      const wt = new WebTTY(term, factory, args, gotty_auth_token);
      const closer = wt.open();

      window.addEventListener("unload", () => {
        closer();
        term.close();
      });
    } else {
      console.log("ele is null")
    }
  }

  render() {
    return (
      <div id="terminal" />
    )
  }
}
