import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { File } from '../../domain/file'
@Injectable()
export class FilesProvider {
  private _days = [
    '2018-08-04',
    '2018-08-03',
    '2018-08-02',
    '2018-08-01',
    '2018-07-31',
    '2017-07-30',
    '2017-07-29',
    '2017-07-28',
    '2017-07-21',
    '2017-07-11',
    '2017-07-10',
  ]
  files = new Map<string, Array<File>>()
  days = []
  constructor(public http: HttpClient) {
    for (const d of this._days.slice(0, 5)) {
      this.getFiles(d).then(files => {
        this.files.set(d, files)
        this.days.push(d)
      })
    }
    this.days.push(...this._days.slice(0, 5))
  }
  getFiles(day: string): Promise<Array<File>> {
    return new Promise<Array<File>>(resolve => {
      resolve([
        {
          id: 'xx',
          type: 'image',
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
      ])
    })
  }
  loadDay(): Promise<boolean> {
    return new Promise<boolean>(resolve => {
      if (this.days.length < this._days.length) {
        const d = this._days[this.days.length]
        this.getFiles(d).then(files => {
          this.files.set(d, files)
          this.days.push(d)
          resolve(true)
        })
      }
      resolve(true)
    })
  }
}
