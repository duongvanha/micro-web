import AxiosProvider from './AxiosProvider';

const axiosProvider = new AxiosProvider({
    BASE_URL: process.env.BASE_URL,
});

export { axiosProvider }