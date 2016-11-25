-----
**Arch List**

-----
*data structure*

 - $Item
  - $digital item
    - $account
    - $record data
  - $real item
    - $room
    - $electronic equipment
    - $useable material
    - ...



    //data structure will be implemented by class extends.
    //item will be used as list management.
    

-----
*relations*

 - $whereis : $id1 is [in/on/..] $id2
 - $connectwith: $id1 is [partof/usedfor/...] $id2



    //relations will be used for search and list



-----
*notice*

 - $starttime: $id1 is start from [$cur] time.
 - $periodtime: $id1 needs notice since $start time.
 - $periodtype: $id1 needs to notice in [once/every/...] $periodtime

 
-----
*modules*

 - $piccontent: pic class to hold pic for $id in filesystem.
 - $notice:  notice module for manage notice thread.
 - $index: relaton management for index list.
 - $item: item management for [add/remove/modify/...]







