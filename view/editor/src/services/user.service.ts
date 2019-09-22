import HttpService from './http.service';

/**
 * Service to manage connection and user
 */
class CUserService {
  private token:string;
  private user:any;

  constructor(){
    this.token = '';
  }

  public isConnected():boolean{
    return false;
  }

}

const UserService = new CUserService();
export default UserService; 