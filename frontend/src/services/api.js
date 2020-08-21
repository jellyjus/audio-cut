const baseURL = process.env.NODE_ENV === 'development' ?
    'http://localhost:3000/api' :
    `${location.origin}/api`;

const getBody = {
    method: 'get',
    credentials: 'include',
    headers: {
        "Content-type": "application/json; charset=utf-8"
    }
};
const postBody = {
    method: 'post',
    credentials: 'include',
    headers: {
        "Content-type": "application/json; charset=utf-8"
    }
};

export default {
    install(Vue, store, router) {
        Vue.prototype.$api = {
            async login() {
                const res = await get('/login');
                if (res.need_redirect)
                    return window.location.replace(res.redirect_url);

                return res
            },
            async getAudios() {
                return get('/get_audios', )
            }
        };
    }
};

async function get(path, params) {
    const url = new URL(baseURL + path);
    url.search = _serializeQuery(params);
    return _send(url, getBody)
}

async function post(path, params) {
    const url = new URL(baseURL + path);
    return _send(url, {...postBody, body: JSON.stringify(params)})
}

async function _send(url, body) {
    let data;
    const res = await fetch(url.toString(), body);

    try {
        data = await res.json();
    } catch (e) {
        data = `Unexpected error. See Network tab in developer console`
    }

    if (!res.ok)
        throw new Error(`${url.pathname} завершился со статусом ${res.status}. ${data.message}`);
    return data
}

function _serializeQuery(obj) {
    const str = [];
    for (const p in obj)
        if (obj[p] != null)
            str.push(encodeURIComponent(p) + "=" + encodeURIComponent(obj[p]));

    return str.join("&");
}

