import AxiosProvider   from './AxiosProvider';
import PhimMoiProvider from './Movie/PhimMoi';

const axiosProvider = new AxiosProvider({
    BASE_URL: process.env.BASE_URL,
});

const phimMoiProvider = new PhimMoiProvider(axiosProvider);

export { axiosProvider, phimMoiProvider }
