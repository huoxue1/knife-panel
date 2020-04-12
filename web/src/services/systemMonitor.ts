import request from '@/utils/request';

export async function basicInfo() {
  return request('/v1/system-monitor', {
  });
}

