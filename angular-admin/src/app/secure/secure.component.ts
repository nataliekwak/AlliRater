import { Component, OnInit } from '@angular/core';
import {AuthService} from '../services/auth.service';
import {Router} from '@angular/router';
import {User} from'../interfaces/user';
import { Star } from '../interfaces/star';

@Component({
  selector: 'app-secure',
  templateUrl: './secure.component.html',
  styleUrls: ['./secure.component.css']
})
export class SecureComponent implements OnInit{
  user: User;
  star: Star;
  stars: Star[];
  constructor(
    private authService: AuthService,
    private router: Router
    ){ }

  ngOnInit(): void{
    this.authService.user().subscribe(
      user => this.user = user,
      err => this.router.navigate(['/login'])
    );
  }

  selectStar(id: any) {
    id = 0
  }

}
