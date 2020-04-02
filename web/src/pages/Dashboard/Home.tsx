import {Card} from 'antd';
import {AnyAction, Dispatch} from 'redux';
import React, {PureComponent} from 'react';

import {ConnectState} from '@/models/connect';
import {GlobalModelState} from '@/models/global';
import PageHeaderLayout from '@/layouts/PageHeaderLayout';
import {connect} from 'dva';
import moment from 'moment';


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
  }

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

  render() {
    const {
      global: {user},
    } = this.props;

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

        </Card>
      </PageHeaderLayout>
    );
  }
}

export default Home;
