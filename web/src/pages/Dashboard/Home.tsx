import {Card, Col, Row} from 'antd';
import {AnyAction, Dispatch} from 'redux';
import React, {PureComponent} from 'react';

import {ConnectState} from '@/models/connect';
import {GlobalModelState} from '@/models/global';
import PageHeaderLayout from '@/layouts/PageHeaderLayout';
import {connect} from 'dva';
import moment from 'moment';
import styles from './Home.less'
import WaterWave from "@/components/Charts/WaterWave";


export interface HomeProps {
  dispatch: Dispatch<AnyAction>;
  global: GlobalModelState;
}

export interface HomeState {
  currentTime: string;
}

@connect(({global}: ConnectState) => ({
  global,
}))
class Home extends PureComponent<HomeProps, HomeState> {
  private interval: NodeJS.Timeout | undefined;

  constructor(props: HomeProps) {
    super(props);
    this.state = {
      currentTime: moment().format('HH:mm:ss'),
    };
  }

  componentDidMount(): void {
    this.interval = setInterval(() => {
      this.setState({currentTime: moment().format('HH:mm:ss')});
    }, 1000);
    this.dispatch({
      type: 'global/fetchBasicSystemInfo',
    });
  }

  dispatch = (action: any) => {
    const {dispatch} = this.props;
    dispatch(action);
  };

  componentWillUnmount(): void {
    if (this.interval !== undefined) {
      clearInterval(this.interval);
    }
  }

  getHeaderContent = () => {
    const {
      global: {user},
    } = this.props;

    const roleNames = user && user.role_names;

    const text = [];
    if (roleNames && roleNames.length > 0) {
      text.push(
        <span key="role" style={{marginRight: 20}}>{`所属角色：${roleNames.join('/')}`}</span>,
      );
    }

    if (text.length > 0) {
      return text;
    }
    return null;
  };

  format(num: number): number {
    if (num == undefined) {
      return 0;
    }
    return Number(num.toFixed(0))
  }

  render() {
    const {
      global: {user, monitor},
    } = this.props;

    const {info_stat, v_mem_stat} = monitor || {}
    const {currentTime} = this.state;
    const breadcrumbList = [{title: '首页'}];

    return (
      <PageHeaderLayout
        title={`您好，${user && user.real_name}，祝您开心每一天！`}
        breadcrumbList={breadcrumbList}
        content={this.getHeaderContent()}
        action={<span>当前时间：{currentTime}</span>}
      >
        <Card>
          <Row>
            <Col span={6} className={styles.dashCol}>
              <Card hoverable className={styles.dashMonitorItem}>
                <div className={styles.systemInfo}>
                  <div><span className={styles.itemLabel}>主机名</span><span>{info_stat.hostname}</span></div>
                  <div><span className={styles.itemLabel}>操作系统</span><span>{info_stat.os}</span></div>
                  <div>
              <span
                className={styles.itemLabel}>平台版本</span><span>{`${info_stat.platform} ${info_stat.platformFamily} ${info_stat.platformVersion}`}</span>
                  </div>
                  <div><span className={styles.itemLabel}>内核版本</span><span>{info_stat.kernelVersion}</span></div>
                </div>
              </Card>
            </Col>
            <Col span={6} className={styles.dashCol}>
              <Card hoverable className={styles.dashMonitorItem}>
                <WaterWave
                  height={64}
                  title={"内存"}
                  percent={this.format(v_mem_stat.usedPercent)}
                />
              </Card>
            </Col>
            <Col span={6} className={styles.dashCol}></Col>
            <Col span={6} className={styles.dashCol}></Col>
          </Row>
        </Card>
      </PageHeaderLayout>
    );
  }
}

export default Home;
