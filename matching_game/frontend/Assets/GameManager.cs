using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class GameManager : MonoBehaviour
{
   [SerializeField] 
    private Camera camera_object; //カメラを取得
    private RaycastHit hit; //レイキャストが当たったものを取得する入れ物

    // Start is called before the first frame update
    void Start()
    {
        for (int i = 1; i <= 16; i++) {

        }
    }

    // Update is called once per frame
    void Update()
    {
      if (Input.GetMouseButtonDown(0)) //マウスがクリックされたら
       {
           Ray ray = camera_object.ScreenPointToRay(Input.mousePosition); //マウスのポジションを取得してRayに代入

               if(Physics.Raycast(ray,out hit))  //マウスのポジションからRayを投げて何かに当たったらhitに入れる
           {
              string objectName = hit.collider.gameObject.name; //オブジェクト名を取得して変数に入れる
               Debug.Log(objectName); //オブジェクト名をコンソールに表示
           }
       }

    }


}
