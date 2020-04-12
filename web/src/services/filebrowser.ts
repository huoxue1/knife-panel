import request from '@/utils/request';

export async function list(params?: any) {
  return request('/v1/file-browser', {
    params,
  });
}

