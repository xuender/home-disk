import { Component, Input, OnInit } from '@angular/core';
import { File } from '../../domain/file'
@Component({
  selector: 'hd-files',
  templateUrl: 'hd-files.html'
})
export class HdFilesComponent implements OnInit {
  @Input() day: string
  files: File[] = []
  constructor() {
  }
  ngOnInit() {
    console.log('day', this.day)
    this.files = [
      {
        id: 'xx',
        type: 'archive',
      },
      {
        id: 'xx',
        type: 'audio',
      },
      {
        id: 'xx',
        type: 'archive',
      },
      {
        id: 'xx',
        type: 'audio',
      },
      {
        id: 'xx',
        type: 'video',
      },
      {
        id: 'xx',
        type: 'video',
      },
      {
        id: 'xx',
        type: 'video',
      },
    ]
  }
}
