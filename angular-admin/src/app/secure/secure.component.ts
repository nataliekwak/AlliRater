import { Component, OnInit } from '@angular/core';
import {AuthService} from '../services/auth.service';
import {Router} from '@angular/router';
import {User} from'../interfaces/user';
import { FormsModule } from '@angular/forms';


@Component({
  selector: 'app-secure',
  templateUrl: './secure.component.html',
  styleUrls: ['./secure.component.css']
})
export class SecureComponent implements OnInit{
  user: User;
  avg: number;
  
  constructor(
    private authService: AuthService,
    private router: Router
    ){ }

    templateForm(value: any) {
      console.log(value);
      alert(JSON.stringify(value));
    }

  ngOnInit(): void{
    this.authService.user().subscribe(
      user => this.user = user,
      err => this.router.navigate(['/login'])
    );
  }

  

}
