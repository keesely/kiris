<?php
/**
 * 
 * @fileName config.php
 * @category PHP
 * @package void
 * @author Kee Guo <chinboy2012@gmail.com> 
 * @since 26/10/2019
 * @version config.php 2019.10.26
 * */

class Conf {
  protected $cnf = [
    "app" => [
      "name" => "kiris",
      "redis" => [
        "host" => "localhost",
        "port" => 6379
      ]
    ]
  ];

  function get ($key, $def = null) {
    if (empty($key)) return $cnf;

    $keys = explode(".", $key);

    $data = $this->cnf;
    foreach ($keys as $k) {
      if (isset($data[$k])) {
        $data = $data[$k];
      }
      else {
        return NULL;
      }
    }
    return $data;
  }

  function set ($key, $value) {
    if (empty($key)) return NULL;
    $keys = explode(".", $key);

    $data = &$this->cnf;
    foreach ($keys as $k) {
      if (isset($data[$k])) {
        $data = &$data[$k];
      }
      else {
        $data[$k] = [];
        $data = &$data[$k];
      }
    }
    $data = $value;
  }
}


$cnf = new Conf();
var_dump($cnf->set("app.redis.host", '127.0.0.1'));
var_dump($cnf->get("app.redis.host"));
