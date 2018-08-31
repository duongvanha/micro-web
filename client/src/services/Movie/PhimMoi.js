const decodeUrlInfoMovie = Symbol();
const encodeStep1        = Symbol();
const encodeStep2        = Symbol();
const encodeStep3        = Symbol();

/**
 *
 * @class
 * @implements {MovieProvider}
 */
export default class PhimMoiProvider {

    constructor(httpService) {
        this.httpService = httpService;
    }

    getMovieByUrl(url) {
        return this[encodeStep1](url)
            // .then(response => this[encodeStep2](response))
            // .then(response => this[decodeUrlInfoMovie](response))
    }

    /**
     *
     * @param {string} stringEncode
     * @returns {{}} movie
     */
    [decodeUrlInfoMovie](stringEncode) {
        return JSON.parse(stringEncode.match(/_responseJson='(.*)'/)[1])
    }

    /**
     *
     * @param url {string}
     * @returns {Promise<string>}
     */
    async [encodeStep1](url) {
        this.httpService.request.get(url).then(console.log).catch(console.log)
        // return requestWithHeader(url)
        //     .then(response => response.match(/(;eval.*'\)\);)/)[1])
    }

    [encodeStep3](w, i, s, e) {
        let lIll = 0;
        let ll1I = 0;
        let Il1l = 0;
        let ll1l = [];
        let l1lI = [];
        while (true) {
            if (lIll < 5) l1lI.push(w.charAt(lIll)); else if (lIll < w.length) ll1l.push(w.charAt(lIll));
            lIll++;
            if (ll1I < 5) l1lI.push(i.charAt(ll1I)); else if (ll1I < i.length) ll1l.push(i.charAt(ll1I));
            ll1I++;
            if (Il1l < 5) l1lI.push(s.charAt(Il1l)); else if (Il1l < s.length) ll1l.push(s.charAt(Il1l));
            Il1l++;
            if (w.length + i.length + s.length + e.length === ll1l.length + l1lI.length + e.length) break;
        }
        let lI1l = ll1l.join('');
        let I1lI = l1lI.join('');
        ll1I     = 0;
        let l1ll = [];
        for (lIll = 0; lIll < ll1l.length; lIll += 2) {
            let ll11 = -1;
            if (I1lI.charCodeAt(ll1I) % 2) ll11 = 1;
            l1ll.push(String.fromCharCode(parseInt(lI1l.substr(lIll, 2), 36) - ll11));
            ll1I++;
            if (ll1I >= l1lI.length) ll1I = 0;
        }
        return l1ll.join('');
    }

    /**
     *
     * @param response {string}
     * @returns {Promise<*>}
     */
    [encodeStep2](response) {
        try {
            const [, p1, p2, p3, p4] = response.match(/\('(\w+)','(\w+)','(\w+)','(\w+)'\)/);
            return this[encodeStep2](this[encodeStep3](p1, p2, p3, p4))
        } catch (e) {
            // return requestWithHeader(response.match(/"(http:.*)" /)[1]);
        }
    }

};
