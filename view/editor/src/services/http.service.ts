import axios from 'axios';

/**
 * Override axios
 */
class CHttpService {

    private serverURL:string = process.env.VUE_APP_SERVER_URL;
    private serverPort:string = process.env.VUE_APP_SERVER_PORT;

    private getURL():string{ return this.serverURL + ':' + this.serverPort + '/'; }

    public getDefaultHeaders():any {
        return {'Content-Type':'application/json'};
    }

    public appendHeaders(headers:any, key:string, value:string):any {
        headers[key] = value;
        return headers;
    }

    public getDisplayServerURL():string{
        var result =  this.serverURL;
        if(this.serverURL == null || this.serverURL.length == 0) result = 'empty server, check your env....';
        else if(this.serverPort == null || this.serverPort.length == 0) result = 'empty port, check your env....';
        else if(this.serverPort.length > 0 && this.serverPort != '80' && this.serverPort != '443'){
            result += ':' + this.serverPort;
        }
        return result;
    }

    public get(url:string, headers:any, callbackSuccess:(response: any) => void, callbackError:(error: any) => void):void {
        axios.get(this.getURL() + url, {headers: headers}).then(callbackSuccess).catch(callbackError);   
    }

    public post(url:string, data:any, headers:any, callbackSuccess:(response: any) => void, callbackError:(error: any) => void):void {
        axios.post(this.getURL() + url, data, {headers: headers}).then(callbackSuccess).catch(callbackError);   
    }

    public put(url:string, data:any, headers:any, callbackSuccess:(response: any) => void, callbackError:(error: any) => void):void {
        axios.put(this.getURL() + url, data, {headers: headers}).then(callbackSuccess).catch(callbackError);   
    }

    public delete(url:string, headers:any, callbackSuccess:(response: any) => void, callbackError:(error: any) => void):void {
        axios.delete(this.getURL() + url, {headers: headers}).then(callbackSuccess).catch(callbackError);   
    }
}

const HttpService = new CHttpService();
export default HttpService; 