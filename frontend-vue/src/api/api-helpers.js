import config from './api.config';

export const formatRequest = request => {
  if (!request) {
    return {};
  }

  return JSON.parse(JSON.stringify(request));
};

export const getHeaders = () => {
  let headers = formatRequest(config.axiosConfig().headers);

  // TODO: add custom headers

  return headers;
};
