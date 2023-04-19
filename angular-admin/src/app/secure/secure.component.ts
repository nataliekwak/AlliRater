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
  ngOnInit(): void {
      
  }

}
