import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/services/auth.service';
import {Router} from '@angular/router';
import { User } from 'src/app/interfaces/user';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit{
  user: User;
  value: string;
  value1: string;
  value2: string;
  num: number;
  sum = 0;
  count = 0;
  sum2 = 0;
  count2 = 0;
  constructor(
    private authService: AuthService,
    private router: Router
    ){ }

    templateForm(value: string) {
        console.log(value);
        //alert(JSON.stringify(value));
        this.value = JSON.stringify(value);
        this.value = this.value.substring(9,10)
        this.sum += Number(this.value)
        this.count += 1;
        this.value = JSON.stringify(this.sum/this.count)
        //console.log(value[0])
    }

    templateForm2(value: string) {
      console.log(value);
      //alert(JSON.stringify(value));
      this.value2 = JSON.stringify(value);
      this.value2 = this.value2.substring(10,11)
      this.sum2 += Number(this.value2)
      this.count2 += 1;
      this.value2 = JSON.stringify(this.sum2/this.count2)
      //console.log(value[0])
    }

  ngOnInit(): void{
    this.authService.user().subscribe(
      user => this.user = user,
      err => this.router.navigate(['/login'])
    );
  }
}
