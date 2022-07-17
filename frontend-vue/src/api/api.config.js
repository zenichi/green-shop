import runtimeConfig from '@/core/runtimeConfig';

let config = {
  apiHost: () => {
    return `${runtimeConfig.apiHost}`;
  },
  axiosConfig: () => {
    return {
      headers: {
        Accept: 'text/plain',
        'Access-Control-Allow-Origin': '*'
      }
    };
  }
};

export default config;
