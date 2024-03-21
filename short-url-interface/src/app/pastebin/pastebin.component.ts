import { Component, OnInit } from '@angular/core';
import { PastebinService } from '../services/pastebin.service';
import { Link, Result, ShortURL } from '../models/model';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-pastebin',
  standalone:true,
  templateUrl: './pastebin.component.html',
  styleUrls: ['./pastebin.component.scss'],
  imports:[CommonModule,FormsModule]
})
export default class PastebinComponent implements OnInit {

  public countLink: number = 0;
  public linkList: Link[] = [];
  public newLink: string ='';
  public shortURL: string = '';

  constructor(private pastebinService: PastebinService) { }

  ngOnInit() {
    this.loadAlls();
  }

  public redirect(key:string) {
    this.pastebinService.redirect(key).subscribe(() => {
    });
  }

  public shorten(urlString:string) {
    this.pastebinService.shorten({ url: urlString }).subscribe(shortURL => {
      this.shortURL = shortURL.url;
      this.loadAlls();

    });
  }

  public loadAlls() {
    this.pastebinService.showAll().subscribe((result:Result) => {
      this.linkList = result?.links;
      this.countLink = result?.total_links;
    });
  }

  public stats(key:string) {
    this.pastebinService.stats(key).subscribe(stats => {
      console.log(stats);
      const link = this.linkList.find(link => link.key === key);
      if (link) {
        link.click = stats.count;
      }
    });
  }
}