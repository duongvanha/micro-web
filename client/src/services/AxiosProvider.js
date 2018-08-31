import Axios from 'axios';

const a = Symbol();

export default class AxiosProvider {

    constructor(config) {
        this[a] = Axios.create(config);
    }

    get request() {
        return this[a];
    }
}
