// Copyright © 2018 Difan Chen <dfchen6@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"encoding/binary"

	"github.com/spf13/cobra"
	"github.com/boltdb/bolt"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		addTask := cmd.Flag("task").Value.String()
		fmt.Printf("%s is added\n", addTask)
		add_task(addTask)
	},
}

func add_task(taskName string) {
	db, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		fmt.Errorf("open bolt: %s", err)
	}

	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TaskBucket"))
		id64, _ := b.NextSequence()
		id := int(id64)
		key := itob(id)
		err := b.Put(key, []byte(taskName))
		
		if err != nil {
			return err
		}
		return nil
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
