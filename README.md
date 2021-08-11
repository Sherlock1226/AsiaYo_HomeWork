# AsiaYo_HomeWork

1.營運人員希望得知最多房間的旅宿前10名的Sql

    SELECT p.`name`
    FROM room AS r
    LEFT JOIN property AS p
    ON (r.`property_id` = p.`id`)
    GROUP BY property_id  
    ORDER BY count( * )  DESC  
    LIMIT 10;
    
    
2.  <img width="499" alt="截圖 2021-08-11 下午2 14 44" src="https://user-images.githubusercontent.com/31307327/128978817-71c52825-bd0a-4c9d-9b75-1261b9bf3e16.png">
  
  使用GOLANG作答
  代碼在tasks
