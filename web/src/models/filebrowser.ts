import * as fileBrowserService from '@/services/filebrowser';

import {Effect} from 'dva';
import {Reducer} from 'redux';

export interface FileBrowserModelState {
  search?: any;
  data?: {
    list: any;
  };
}

export interface FileBrowserModelType {
  namespace: 'fileBrowser';
  state: FileBrowserModelState;
  effects: {
    fetch: Effect;
  };
  reducers: {
    saveData: Reducer<FileBrowserModelState>;
  };
}

const FileBrowserModel: FileBrowserModelType = {
  namespace: 'fileBrowser',
  state: {
    search: {},
    data: {
      list: [],
    },
  },
  effects: {
    * fetch({search, pagination}, {call, put, select}) {
      const response = yield call(fileBrowserService.list, search);
      yield put({
        type: 'saveData',
        payload: {list: response},
      });
    },

  },
  reducers: {
    saveData(state, {payload}) {
      return {...state, data: payload};
    },
  },
};

export default FileBrowserModel;
