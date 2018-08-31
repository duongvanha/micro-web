import Axios from 'axios';

const a = Symbol();

export default class AxiosProvider {

    constructor() {
        this[a] = Axios.create({
            headers: {'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8'},
        });
    }

    get request() {
        return this[a]
    }
}
