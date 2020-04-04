import {Form} from '@ant-design/compatible';
import '@ant-design/compatible/assets/index.css';
import {Table} from 'antd';
import React from 'react';
import {FormComponentProps} from '@ant-design/compatible/es/form';

import {AnyAction, Dispatch} from "redux";
import {RoleModeState} from "@/models/role";
import {connect} from "dva";
import {ConnectState} from "@/models/connect";

export interface FileBrowserProps extends FormComponentProps {
  dispatch: Dispatch<AnyAction>;
  loading: any;
  fileItem: RoleModeState;
}

export interface FileBrowserState {
}

@connect((state: ConnectState) => ({
  role: state.role,
  loading: state.loading.models.role,
}))
class FileBrowser extends React.PureComponent<FileBrowserProps, FileBrowserState> {

  // dispatch = (action: any) => {
  //   const {dispatch} = this.props;
  //   dispatch(action);
  // };
  //
  // componentDidMount() {
  //   this.dispatch({
  //     type: 'role/fetch',
  //     search: {},
  //     pagination: {},
  //   });
  // }

  render() {
    return (<div></div>)
  //   const {
  //     fileItem: {data},
  //   } = this.props;
  //
  //   const {list} = data || {};
  //   const columns = [
  //     {
  //       title: '名称',
  //       dataIndex: 'name',
  //     },
  //     {
  //       title: '修改时间',
  //       dataIndex: 'modifyTime',
  //       sorter: true,
  //       valueType: 'dateTime',
  //     },
  //   ];
  //
  //   return (
  //     <div>
  //       <Table columns={columns} dataSource={list}/>
  //     </div>
  //   );
  }

}

export default Form.create<FileBrowserProps>()(FileBrowser);
