// Solo.go - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cache

import (
	"github.com/b3log/solo.go/model"
	"github.com/bluele/gcache"
)

var Comment = &commentCache{
	idHolder: gcache.New(1024).LRU().Build(),
}

type commentCache struct {
	idHolder gcache.Cache
}

func (cache *commentCache) Put(comment *model.Comment) {
	cache.idHolder.Set(comment.ID, comment)
}
