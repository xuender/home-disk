import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { File } from '../../domain/file'
import { Observable } from 'rxjs/Observable';
@Injectable()
export class FilesProvider {
  private _days: string[] = []
  files = new Map<string, Array<File>>()
  days = []
  private run = true
  constructor(public http: HttpClient) {
    this.http.get('/days')
      .subscribe((days: string[]) => {
        this._days = days
        // 初始化 5 天
        for (let i = 0; i < 5; i++) {
          const d = this._days.shift()
          if (d) {
            console.log('d', d)
            this.getFiles(d).subscribe(files => {
              this.files.set(d, files)
              this.days.push(d)
            })
          } else {
            break
          }
        }
        this.run = false
      })
  }
  getFiles(day: string): Observable<Array<File>> {
    return this.http.get<Array<File>>(`/days/${day}`)
  }
  loadDay(): Promise<boolean> {
    console.log('run', this.run)
    return new Promise<boolean>(resolve => {
      if (!this.run && this._days.length > 0) {
        this.run = true
        const d = this._days.shift()
        if (d) {
          console.log('dd', d)
          this.getFiles(d).subscribe(files => {
            this.files.set(d, files)
            this.days.push(d)
            this.run = false
            resolve(true)
          })
        } else {
          this.run = false
        }
      }
      resolve(true)
    })
  }
}
